export const config = {
  port: parseInt(process.env["HOST_PORT"] ?? "4000"),
  dbHost: process.env["PGHOST"] ?? "",
  dbPassword: process.env["PGPASSWORD"] ?? "",
  dbUser: process.env["PGUSER"] ?? "",
  dbPort: parseInt(process.env["PGPORT"] ?? "5432"),
};
