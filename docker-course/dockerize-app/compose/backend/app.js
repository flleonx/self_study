import express from 'express';
import cors from 'cors';

const app = express();

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/', (_req, res) => {
  console.log('Receiving request');
  return res.json('Server');
});

const PORT = 59000;

app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`);
});
