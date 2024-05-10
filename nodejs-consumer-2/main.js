const { Kafka, logLevel } = require('kafkajs');
require('dotenv').config();

const kafkaBrokers = process.env.KAFKA_BROKERS;
const username = process.env.KAFKA_API_KEY;
const password = process.env.KAFKA_API_SECRET;

const kafka = new Kafka({
  logLevel: logLevel.INFO,
  brokers: [kafkaBrokers],
  ssl: {
    rejectUnauthorized: true,
  },
  sasl: {
    mechanism: 'plain',
    username,
    password,
  },
});

const topic = 'order';
const consumer = kafka.consumer({ groupId: 'NCM-test' });

const run = async () => {
  await consumer.connect();
  await consumer.subscribe({ topic, fromBeginning: true });
  await consumer.run({
    // eachBatch: async ({ batch }) => {
    //   console.log(batch)
    // },
    eachMessage: async ({ topic, partition, message }) => {
      const prefix = `2-${topic}[${partition} | ${message.offset}] / ${message.timestamp}`;
      console.log(`- ${prefix} ${message.key}#${message.value}`);
    },
  });
};

run().catch((e) => console.error(`[example/consumer] ${e.message}`, e));

const errorTypes = ['unhandledRejection', 'uncaughtException'];
const signalTraps = ['SIGTERM', 'SIGINT', 'SIGUSR2'];

errorTypes.forEach((type) => {
  process.on(type, async (e) => {
    try {
      console.log(`process.on ${type}`);
      console.error(e);
      await consumer.disconnect();
      process.exit(0);
    } catch (_) {
      process.exit(1);
    }
  });
});

signalTraps.forEach((type) => {
  process.once(type, async () => {
    try {
      await consumer.disconnect();
    } finally {
      process.kill(process.pid, type);
    }
  });
});
