import * as cloudamqp from "@pulumi/cloudamqp";

const instance = new cloudamqp.Instance("my-typescript-instance", {
    plan: "lemur",
    region: "amazon-web-services::us-west-2",
    rmqVersion: "3.12.13",
});

export const name = instance.name;
