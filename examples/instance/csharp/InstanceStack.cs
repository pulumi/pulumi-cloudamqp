using Pulumi;
using CloudAmqp = Pulumi.CloudAmqp;

class InstanceStack : Stack
{
    public InstanceStack()
    {
        var instance = new CloudAmqp.Instance("my-csharp-instance", new CloudAmqp.InstanceArgs
        {
            Plan = "lemur",
            Region = "amazon-web-services::us-west-2",
            RmqVersion = "3.12.13",
        });
    }
}
