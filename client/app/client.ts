import { HappenedService } from "@/gen/protos/protos/v1/happened_service_pb";
import { createXHRGrpcWebTransport } from "./custom-transport";
import { createClient } from "@connectrpc/connect";

export const client = createClient(
  HappenedService,
  createXHRGrpcWebTransport({
    baseUrl: "http://localhost:8080",
  }),
);
