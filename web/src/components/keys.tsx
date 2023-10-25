import { useQuery } from "@tanstack/react-query";
// import { protoInt64 } from "@bufbuild/protobuf";
import { Params } from "react-router-dom";
import { useLoaderData } from "react-router";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Key as KeyProto } from "../proto/fusionchain/treasury/key_pb";
import { keyRequests, keys } from "../client/treasury";
import { prettyBytes, prettyKeyType } from "../utils/formatting";
import { Link } from "react-router-dom";
// import { MsgNewWalletRequest } from "../proto/fusionchain/treasury/tx_pb";
import { WalletType } from "../proto/fusionchain/treasury/wallet_pb";
import { useKeplrAddress } from "../keplr";
import { Button } from "./ui/button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { useBroadcaster } from "@/hooks/keplr";
import { workspaceByAddress } from "@/client/identity";
import { KeyResponse, WalletKeyResponse } from "@/proto/fusionchain/treasury/query_pb";

export default function Keys({ workspaceAddr }: { workspaceAddr: string }) {
  const wsQuery = useQuery({ queryKey: ["keys"], queryFn: () => keys(workspaceAddr) });
  //console.log(wsQuery)

  return (
    <div className="p-4 space-y-3">
      {wsQuery.data?.keys.map((key) => <Key key={key.key?.id.toString()} keyData={key.key!} />)}
    </div>
  );
}

function Wallets({ workspaceAddr }: { workspaceAddr: string }){
  const walletsQuery = useQuery({ 
    queryKey: ["keys", workspaceAddr], 
    queryFn: () => keys(workspaceAddr, 2)
  });

  console.log(walletsQuery)

  return (
    <div className="p-4 space-y-3">
      {walletsQuery.data?.keys.map((walletresponse) => (
        <WalletResponse key={walletresponse.wallets} keyData={walletresponse.wallets}/>
      ))}
    </div>
  );
}

function Key({ keyData }: { keyData: KeyProto }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Key #{keyData.id.toString()}{" "}</CardTitle>
        <CardDescription>Managed by <KeyringLink keyringAddr={keyData.keyringAddr} />.</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid w-full items-center gap-4">
          <div className="flex flex-col space-y-1">
            <span className="text-sm font-bold">Type</span>
            <span>{prettyKeyType(keyData.type)}</span>
          </div>
          <div className="flex flex-col space-y-1">
            <span className="text-sm font-bold">Key material</span>
            <span className="font-mono break-all">{prettyBytes(keyData.publicKey)}</span>
          </div>
        </div>
      </CardContent>
      <CardFooter>
        <Link to={`/sign-data/${keyData.id}`}>
          <Button variant="secondary" size="sm">
            Sign arbitrary data
          </Button>
        </Link>
      </CardFooter>
    </Card>
  );
}

function WalletResponse({ keyData }: { keyData: WalletKeyResponse }) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Key #{keyData.id.toString()}{" "}</CardTitle>
        <CardDescription>Managed by <KeyringLink keyringAddr={keyData.keyringAddr} />.</CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid w-full items-center gap-4">
          <div className="flex flex-col space-y-1">
            <span className="text-sm font-bold">Type</span>
            <span>{prettyKeyType(keyData.type)}</span>
          </div>
          <div className="flex flex-col space-y-1">
            <span className="text-sm font-bold">Key material</span>
            <span className="font-mono break-all">{prettyBytes(keyData.publicKey)}</span>
          </div>
        </div>
      </CardContent>
      <CardFooter>
        <Link to={`/sign-data/${keyData.id}`}>
          <Button variant="secondary" size="sm">
            Sign arbitrary data
          </Button>
        </Link>
      </CardFooter>
    </Card>
  );
}

/*
function Wallets({ keyId }: { keyId: bigint }) {
  // const addr = useKeplrAddress();
  // const { broadcast } = useBroadcaster();
  // const walletsQuery = useQuery({ queryKey: ["wallets", keyId.toString()], queryFn: () => wallets(keyId) });
  // const wQuery = useQuery({queryKey: ["wallets", keys("qredoworkspace10j06zdk5gyl6vrss5d5")]});
  const { wallets } = useLoaderData() as Awaited<ReturnType<typeof loader>>;
  const wQuery2 = useQuery({queryKey: ["wallets", keys("qredoworkspace10j06zdk5gyl6vrss5d5", 2)]});
  // const wsQuery3 = useQuery(["workspace", workspaceAddr], () => workspaceByAddress(workspaceAddr));
  const wsQuery4 = useQuery(["wallets", wallets ], () => keys("qredoworkspace10j06zdk5gyl6vrss5d5",2) )


  // const w = walletsQuery.data?.wallets;
  const w2 = wQuery2.data;
  if (!w2) {
    return <div>loading</div>;
  }

  /*
  const possibleWallets = [
    { name: "Ethereum Sepolia", type: WalletType.ETH_SEPOLIA, onClick: () => {
        broadcast([
          new MsgNewWalletRequest({
            creator: addr,
            keyId: protoInt64.parse(keyId),
            walletType: WalletType.ETH_SEPOLIA,
          }),
        ]);
      }
    }
  ]; 
 
  const missingWallets => !w2.find(w => w.wallet?.type === wallet.type));

  return (
    <div className="">
      {
        w.length === 0 && (
          <div className="text-sm text-gray-500">
            This key has no wallets attached to it. Attach a wallet to sign transactions.
          </div>
        )
      }

      {w.map((wallet) => (
        <div key={wallet.wallet?.id.toString()}>
          <span>
            <span className="font-semibold">Ethereum Sepolia: </span>
            <span key={wallet.wallet?.id.toString()} className="font-mono">
              {wallet.address}
            </span>
            <Link to={`/wallet/${wallet.wallet?.id.toString()}`}>
              <Button variant="default" size="sm" className="ml-3">
                Sign transactions
              </Button>
            </Link>
          </span>
        </div>
      ))}

      {missingWallets.length > 0 && (
        <Select onValueChange={(value) => missingWallets.find(w => w.name === value)?.onClick()}>
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Derive a new wallet" />
          </SelectTrigger>
          <SelectContent>
            {missingWallets.map((wallet) => (
              <SelectItem value={wallet.name} key={wallet.name}>{wallet.name}</SelectItem>
            ))}
          </SelectContent>
        </Select>
      )}
    </div>
  );
}
*/

function KeyringLink({ keyringAddr }: { keyringAddr: string }) {
  return (
    <Link className="underline" to={`/keyrings/${keyringAddr}`}>
      Keyring #{keyringAddr}
    </Link>
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
