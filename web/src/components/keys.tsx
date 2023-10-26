import { useQuery } from "@tanstack/react-query";
import { Params } from "react-router-dom";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Key as KeyProto } from "../proto/fusionchain/treasury/key_pb";
import { keys } from "../client/treasury";
import { prettyBytes, prettyKeyType } from "../utils/formatting";
import { Link } from "react-router-dom";
import { Button } from "./ui/button";

export default function Keys({ workspaceAddr }: { workspaceAddr: string }) {
  const wsQuery = useQuery({ queryKey: ["keys"], queryFn: () => keys(workspaceAddr) });
  //console.log(wsQuery)

  return (
    <div className="p-4 space-y-3">
      {wsQuery.data?.keys.map((key) => <Key key={key.key?.id.toString()} keyData={key.key!} />)}
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
