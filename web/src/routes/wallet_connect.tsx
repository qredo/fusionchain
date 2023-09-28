import { useCallback, useEffect, useRef, useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Core } from '@walletconnect/core'
import { RELAYER_EVENTS } from '@walletconnect/core'
import { IWeb3Wallet, Web3Wallet, Web3WalletTypes } from '@walletconnect/web3wallet'
import { buildApprovedNamespaces } from '@walletconnect/utils'
import { ProposalTypes, PendingRequestTypes } from "@walletconnect/types";
import { AuthEngineTypes } from "@walletconnect/auth-client";
import { useToast } from '@/components/ui/use-toast';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

function useWeb3Wallet(relayUrl: string) {
  const [w, setW] = useState<IWeb3Wallet | null>(null);
  const [sessionProposals, setSessionProposals] = useState<ProposalTypes.Struct[]>([]);
  const [authRequests, setAuthRequests] = useState<AuthEngineTypes.PendingRequest[]>([]);
  const [sessionRequests, setSessionRequests] = useState<PendingRequestTypes.Struct[]>([]);

  useEffect(() => {
    if (w) {
      return;
    }

    const core = new Core({
      projectId: "4fda584de3c28e97dfa5847023e337c8",
      relayUrl,
      logger: "debug",
    })

    Web3Wallet.init({
      core,
      metadata: {
        name: 'Fusion Chain wallets',
        description: 'Fusion Chain WalletConnect',
        url: 'https://qredo.com/',
        icons: ['https://avatars.githubusercontent.com/u/37784886'],
      }
    }).then(async wallet => {
      try {
        const clientId = await wallet.engine.signClient.core.crypto.getClientId();
        console.log('WalletConnect ClientID: ', clientId);
        localStorage.setItem('WALLETCONNECT_CLIENT_ID', clientId);
        setW(wallet);
      } catch (error) {
        console.error('Failed to set WalletConnect clientId in localStorage: ', error)
      }
    });

    return () => {
      setW(null);
    };
  }, []);

  const updateState = useCallback(() => {
    if (!w) {
      return;
    }

    setSessionProposals([...w.getPendingSessionProposals() as any as ProposalTypes.Struct[]]);
    setAuthRequests([...w.getPendingAuthRequests() as any as AuthEngineTypes.PendingRequest[]]);
    setSessionRequests([...w.getPendingSessionRequests()]);
  }, [w])

  useEffect(() => {
    if (!w) {
      return;
    }

    w.on('session_proposal', updateState);
    w.on('auth_request', updateState);
    w.on('session_request', updateState);
    w.on('session_delete', updateState);

    // TODOs
    const onSessionPing = (data: any) => console.log('ping', data);
    w.engine.signClient.events.on('session_ping', onSessionPing);

    return () => {
      w.off('session_proposal', updateState);
      w.off('auth_request', updateState);
      w.off('session_request', updateState);
      w.off('session_delete', updateState);
      w.engine.signClient.events.off('session_ping', onSessionPing);
    };
  }, [w]);

  return {
    w,
    sessionProposals,
    authRequests,
    sessionRequests,
  };
}

const supportedNamespaces = {
  eip155: {
    chains: [
      'eip155:1',
      'eip155:5',
      'eip155:11155111',
    ],
    methods: [
      'personal_sign',
      'eth_sign',
      'eth_signTransaction',
      'eth_signTypedData',
      'eth_signTypedData_v3',
      'eth_signTypedData_v4',
      'eth_sendRawTransaction',
      'eth_sendTransaction'
    ],
    events: ['accountsChanged', 'chainChanged'],
    accounts: [
      // fake hardcoded ETH address
      `eip155:1:0x965C8f3C371ef27B4aFD701821Aa3070AbE4b57d`, // mainnet
      `eip155:5:0x965C8f3C371ef27B4aFD701821Aa3070AbE4b57d`, // goerli
      `eip155:11155111:0x965C8f3C371ef27B4aFD701821Aa3070AbE4b57d`, // sepolia
    ],
  },
}

async function approveSession(w: IWeb3Wallet, proposal: any) {
  console.log('approving session proposal', proposal)
  const { id, relays } = proposal;

  const namespaces = buildApprovedNamespaces({
    proposal,
    supportedNamespaces
  })

  console.log('approving namespaces:', namespaces)

  try {
    const session = await w.approveSession({
      id,
      relayProtocol: relays[0].protocol,
      namespaces
    });
    console.log('session proposal approved. Session:', session)
  } catch (e) {
    console.error('Failed to approve session', e)
  }
}

export default function WalletConnectPage() {
  return (
    <div>
      <WalletConnect />
    </div>
  );
}

function WalletConnect() {
  const { w, sessionProposals, authRequests, sessionRequests } = useWeb3Wallet('wss://relay.walletconnect.org');
  const [loading, setLoading] = useState(false)
  const [uri, setUri] = useState("");

  return (
    <>
      <Card>
        <CardHeader>
        </CardHeader>
        <CardContent>
          <form onSubmit={async (e) => {
            e.preventDefault();
            try {
              setLoading(true);
              await w?.pair({ uri });
              console.log('WalletConnect session paired');
            } catch (error) {
              console.error(error);
            } finally {
              setUri('');
              setLoading(false);
            }
          }}>
            <Input type="text" placeholder="Enter WalletConnect URI" value={uri} onChange={e => setUri(e.target.value)} />
            <Button disabled={loading} type="submit">Connect</Button>
          </form >
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Session Proposals</CardTitle>
        </CardHeader>
        <CardContent>
          {sessionProposals.map((proposal, i) => (
            <div key={i}>
              <pre>{JSON.stringify(proposal, null, 2)}</pre>
              <Button disabled={!w} onClick={() => approveSession(w!, proposal)}>Approve</Button>
            </div>
          ))}
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Auth Requests</CardTitle>
        </CardHeader>
        <CardContent>
          {authRequests.map((req, i) => (
            <div key={i}>
              <pre>{JSON.stringify(req, null, 2)}</pre>
              <Button disabled={!w}>Approve</Button>
            </div>
          ))}
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Session Requests</CardTitle>
        </CardHeader>
        <CardContent>
          {sessionRequests.map((req, i) => (
            <div key={i}>
              <pre>{JSON.stringify(req, null, 2)}</pre>
              <Button disabled={!w} onClick={async () => {
                const topic = req.topic;
                let response = null;

                switch (req.params.request.method) {
                  case 'personal_sign': {
                    const msg = req.params.request.params[0];
                    const address = req.params.request.params[1];
                    const signature = "0xdeadbeef";
                    response = {
                      result: signature,
                      id: req.id,
                      jsonrpc: "2.0",
                    };
                  }
                }

                if (!response) {
                  console.error('Unknown session request method', req.params.request.method);
                  return;
                }
                
                await w!.respondSessionRequest({
                  topic,
                  response,
                });
              }}>Approve</Button>
            </div>
          ))}
        </CardContent>
      </Card>
    </>
  )
}
