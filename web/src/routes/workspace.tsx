import { protoInt64 } from "@bufbuild/protobuf";
import { useLoaderData } from "react-router";
import { Params } from "react-router-dom";
import { useKeplrAddress } from "../keplr";
import { keplrBuildAndBroadcast } from "../newclient";
import { MsgNewKeyRequest } from "../proto/fusionchain/treasury/tx_pb";
import { KeyType } from "../proto/fusionchain/treasury/key_pb";
import Keys from "../components/keys";
import KeyRequests from "../components/key_requests";
import { workspaceByAddress } from "../client/identity";
import { useQuery } from "@tanstack/react-query";
import Address from "../components/address";
import { MsgAddWorkspaceOwner, MsgRemoveWorkspaceOwner } from "../proto/fusionchain/identity/tx_pb";
import { useState } from "react";

async function requestNewKey(creator: string, workspaceAddr: string, keyringId: number) {
  await keplrBuildAndBroadcast([
    new MsgNewKeyRequest({ keyringId: protoInt64.parse(keyringId), creator, workspaceAddr, keyType: KeyType.ECDSA_SECP256K1 }),
  ]);
}

async function addOwner(creator: string, workspaceAddr: string, newOwner: string) {
  await keplrBuildAndBroadcast([
    new MsgAddWorkspaceOwner({ creator, workspaceAddr, newOwner }),
  ]);
}

async function removeOwner(creator: string, workspaceAddr: string, owner: string) {
  await keplrBuildAndBroadcast([
    new MsgRemoveWorkspaceOwner({ creator, workspaceAddr, owner }),
  ]);
}

function Workspace() {
  const addr = useKeplrAddress();
  const { workspaceAddr } = useLoaderData() as Awaited<ReturnType<typeof loader>>;
  const wsQuery = useQuery(["workspace", workspaceAddr], () => workspaceByAddress(workspaceAddr));
  const [newOwner, setNewOwner] = useState("");
  const keyringId = 0;

  return (
    <div className="px-6 mt-10">
      <div className="flex flex-row justify-between items-center">
        <div>
          <h1 className="font-bold text-lg">Workspace {workspaceAddr}</h1>
        </div>
      </div>

      <div className="mt-10 flex flex-col">
        <h2 className="font-bold">Owners:</h2>
        <ul className="list-disc list-inside">
          {
            wsQuery.data?.workspace?.owners.map((owner) => (
              <li key={owner}>
                <Address address={owner} />
                <button className="hover:bg-blue-200 ml-2 px-2 py-1 rounded-lg text-xs" onClick={() => removeOwner(addr, workspaceAddr, owner)}>❌</button>
              </li>
            ))
          }
        </ul>
        <div className="flex flex-row items-center mt-4 gap-2">
          <input className="px-3 py-1 border border-slate-200 rounded-lg" type="text" placeholder="Add new owner" value={newOwner} onChange={(v) => setNewOwner(v.target.value)} />
          <button className="bg-slate-200 hover:bg-blue-200 px-3 py-1 rounded-lg" onClick={async () => {
            await addOwner(addr, workspaceAddr, newOwner);
            setNewOwner("");
          }}>
            Add
          </button>
        </div>
      </div>

      <div className="mt-10 flex flex-row justify-between items-center">
        <div>
          <h2 className="font-bold">Keys</h2>
        </div>

        <button className="bg-slate-200 hover:bg-blue-200 px-4 py-2 rounded-lg" onClick={() => requestNewKey(addr, workspaceAddr, keyringId)}>
          Request a new key
        </button>
      </div>

      <KeyRequests workspaceAddr={workspaceAddr} />

      <Keys workspaceAddr={workspaceAddr} />
    </div>
  );
}

export async function loader({ params }: { params: Params<string> }) {
  if (!params.workspaceAddr) {
    throw new Error("No workspace address provided");
  }
  return {
    workspaceAddr: params.workspaceAddr,
  };
}

export default Workspace;
