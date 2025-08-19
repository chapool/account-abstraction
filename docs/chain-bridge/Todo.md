1. 文档中涉及资产CPOP不要特殊处理，统一当ERC20资产处理
2. 涉及UserOperation批处理的逻辑不需要，因为bundler会自动帮忙处理； 这里批处理需要实现的是我们设计的MasterAggregator功能
3. 应用层批处理功能， 例如：
    - 为用户批量发积分功能，调用CPOPToken的batchMint函数
    - 为用户批量扣积分功能，调用CPOPToken的batchBurn函数