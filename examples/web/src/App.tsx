import './App.css'
import { ConnectError, createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { Library } from "./gen/apis/example/service_connect";
import { useState } from "react";
import { Shelf } from './gen/apis/example/resource_pb';

const baseUrl = document.location.href.replace("5173", "8080");

const transport = createConnectTransport({
  baseUrl: baseUrl,
});

const client = createPromiseClient(Library, transport);

function App() {
  const [inputValue, setInputValue] = useState("");
  const [shelf, setShelf] = useState<Shelf | null>(null)
  const [error, setError] = useState<ConnectError | null>(null);
  return <>
    <form onSubmit={async (e) => {
      e.preventDefault();
      try {
        const response = await client.getShelf({ name: inputValue });
        setShelf(response);
      } catch (err) {
        const connectErr = ConnectError.from(err);
        setError(connectErr);
      }
    }}>
      <p>{baseUrl}</p>
      {shelf && <p>shelf = {shelf.toJsonString()}</p>}
      {error && <p>error = {error.message}</p>}
      <input value={inputValue} onChange={e => setInputValue(e.target.value)} />
      <button type="submit">Send</button>
    </form>
  </>;
}

export default App
