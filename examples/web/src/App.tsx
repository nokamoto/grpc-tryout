import "./App.css";
import { useState } from "react";
import doc from "./gen/example/apis/example/service.pb.json";
import { Method, Proto, Service } from "./gen/apis/tryout/tryout_pb";

const baseUrl = document.location.href.replace("5173", "8080");

function createMethod(method: Method) {
  const [inputValue, setInputValue] = useState<any>({});
  const [shelf, setShelf] = useState<any | null>(null);
  const [error, setError] = useState<any | null>(null);

  return (
    <div>
      <h3>{method.name}</h3>
      {shelf && <p>shelf = {JSON.stringify(shelf)}</p>}
      {error && <p>error = {JSON.stringify(error)}</p>}
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          await fetch(new URL("/example.Library/GetShelf", baseUrl), {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(inputValue),
          }).then((res) => {
            if (res.ok) {
              return res.json().then(setShelf);
            }
            return res.json().then(setError);
          });
        }}
      >
        {method.fields.map((field) => (
          <label>
            {field}
            <input
              value={inputValue[field]}
              onChange={(e) => {
                setInputValue({ ...inputValue, [field]: e.target.value });
              }}
            />
          </label>
        ))}
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
