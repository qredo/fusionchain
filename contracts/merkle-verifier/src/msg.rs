use cosmwasm_schema::{cw_serde, QueryResponses};

#[cw_serde]
pub struct InstantiateMsg {
    pub result: String,
}

#[cw_serde]
pub enum ExecuteMsg {
    Verify {
        proof: String,
        root: Vec<u8>,
        leaf_indices: Vec<u64>,
        leaf_hashes: Vec<Vec<u8>>,
        leaf_count: u64,
    },
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    // GetCount returns the current count as a json-encoded number
    #[returns(GetResultResponse)]
    GetVerifResult {},
}

// We define a custom struct for each query response
#[cw_serde]
pub struct GetResultResponse {
    pub result: String,
}
