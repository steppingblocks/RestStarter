import { Client } from 'pg'
import process from 'process'

const DB_PORT = process.env.PORT ? parseInt(process.env.PORT) : 5432

export const db = new Client({
  user: process.env.PGUSER,
  host: process.env.PGHOST,
  password: process.env.PGPASSWORD,
  port: DB_PORT,
  database: process.env.PGDBNAME,
})
