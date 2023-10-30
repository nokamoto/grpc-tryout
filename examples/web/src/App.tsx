import "./App.css";
import { ConnectError, createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { Library } from "./gen/apis/example/service_connect";
import { SetStateAction, useState } from "react";
import { Shelf } from "./gen/apis/example/resource_pb";
import doc from "./gen/example/apis/example/service.pb.json";
import { Method, Proto, Service } from "./gen/apis/tryout/tryout_pb";
import { GetShelfRequest } from "./gen/apis/example/service_pb";

const baseUrl = document.location.href.replace("5173", "8080");

const transport = createConnectTransport({
  baseUrl: baseUrl,
});

const client = createPromiseClient(Library, transport);

function createField(
  field: string,
  value: GetShelfRequest,
  set: React.Dispatch<SetStateAction<GetShelfRequest>>,
) {
  return (
    <label>
      {field}
      <input
        value={value.name}
        onChange={(e) => {
          const v = value.clone();
          v.name = e.target.value;
          set(v);
        }}
      />
    </label>
  );
}

function createMethod(method: Method) {
  const [inputValue, setInputValue] = useState(new GetShelfRequest());
  const [shelf, setShelf] = useState<Shelf | null>(null);
  const [error, setError] = useState<ConnectError | null>(null);

  return (
    <div>
      <h3>{method.name}</h3>
      {shelf && <p>shelf = {shelf.toJsonString()}</p>}
      {error && <p>error = {error.message}</p>}
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          try {
            const response = await client.getShelf(inputValue);
            setShelf(response);
          } catch (err) {
            const connectErr = ConnectError.from(err);
            setError(connectErr);
          }
        }}
      >
        <label>
          {method.fields.map((field) =>
            createField(field, inputValue, setInputValue),
          )}
        </label>
        <button type="submit">Send</button>
      </form>
    </div>
  );
}

function createService(service: Service) {
  return (
    <div>
      <h2>{service.name}</h2>
      <div>{service.methods.map(createMethod)}</div>
    </div>
  );
}

function App() {
  return <div>{Proto.fromJson(doc).services.map(createService)}</div>;
}

export default App;
