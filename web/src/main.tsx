import React from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import "./main.css";
import Root from "./routes/root.tsx";
import Home from "./routes/home.tsx";
import { QueryClient } from "@tanstack/query-core";
import { QueryClientProvider } from "@tanstack/react-query";
import Workspace, { loader as workspaceLoader } from "./routes/workspace.tsx";
import SignData, { loader as signDataLoader } from "./routes/sign-data.tsx";
import Wallet, { loader as walletLoader } from "./routes/wallet.tsx";
import PoliciesPage from "./routes/policies.tsx";
import ExplorerPage from "./routes/explorer.tsx";
import BlockByHeightPage, { loader as blockByHeightLoader } from "./routes/block_by_height.tsx";
import TxByHashPage, { loader as txByHashLoader } from "./routes/tx_by_hash.tsx";
import ActionsPage from "./routes/actions.tsx";
import WalletConnectPage from "./routes/wallet_connect.tsx";
import KeyringsPage from "./routes/keyrings.tsx";
import Keyring, { loader as keyringLoader } from "./routes/keyring.tsx";

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchInterval: 1000,
    },
  },
});

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "/",
        element: <Home />,
      },
      {
        path: "/policies",
        element: <PoliciesPage />,
      },
      {
        path: "/actions",
        element: <ActionsPage />
      },
      {
        path: "/explorer",
        element: <ExplorerPage />,
      },
      {
        path: "/explorer/block_by_height/:height",
        element: <BlockByHeightPage />,
        loader: blockByHeightLoader,
      },
      {
        path: "/explorer/tx_by_hash/:hash",
        element: <TxByHashPage />,
        loader: txByHashLoader,
      },
      {
        path: "/workspaces/:workspaceAddr",
        element: <Workspace />,
        loader: workspaceLoader,
      },
      {
        path: "/sign-data/:keyId",
        element: <SignData />,
        loader: signDataLoader,
      },

      {
        path: "/wallet/:workspaceAddr/:keyId",
        element: <Wallet />,
        loader: walletLoader,
      },
      {
        path: "/walletconnect",
        element: <WalletConnectPage />,
      },
      {
        path: "/keyrings",
        element: <KeyringsPage />,
      },
      {
        path: "/keyrings/:keyringAddr",
        element: <Keyring />,
        loader: keyringLoader,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>,
);
