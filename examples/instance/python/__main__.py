from pulumi import export
import pulumi_cloudamqp as cloudamqp

instance = cloudamqp.Instance("my-python-instance",
                                plan="lemur",
                                region="amazon-web-services::us-west-2",
                                rmq_version="3.7.18")

export("name", instance.name)
