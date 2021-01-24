# DDB CDC to Kinesis

Inspired by [this announcement](https://aws.amazon.com/about-aws/whats-new/2020/11/now-you-can-use-amazon-kinesis-data-streams-to-capture-item-level-changes-in-your-amazon-dynamodb-table/)

Overall, the ability to stream CDC events from DDB to Kinesis sounds pretty good, especially because we then have the ability to stream to Kinesis Firehose. From there we have a lot of possibilities - eg. ElasticSearch

As always, if something sounds too good to be true, there is some caveat to it. Here it's the fact that you are responsible for sharding the stream, which can be cumbersome. There is an equation you can use to calculate the number of shards though.
