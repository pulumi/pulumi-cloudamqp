import * as cloudamqp from "@pulumi/cloudamqp";

const instance = new cloudamqp.Instance("my-instance", {
    plan: "lemur",
    region: "amazon-web-services::us-west-2"
});

export const name = instance.name;
