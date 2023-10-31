import "./App.css";
import { useState } from "react";
import doc from "./gen/example/apis/example/service.pb.json";
import { Method, Proto, Service } from "./gen/apis/tryout/tryout_pb";

const baseUrl = document.location.href.replace("5173", "8080");

function MethodForm(props: { method: Method }) {
  const initial = () => {
    const v: any = {};
    props.method.fields.forEach((field) => {
      v[field] = "";
    });
    return v;
  };

  const [request, setRequest] = useState<any>(initial());
  const [response, setResponse] = useState<any>({});
  const [error, setError] = useState<any>({});

  return (
    <div>
      <h3>{props.method.name}</h3>
      {response && <p>shelf = {JSON.stringify(response)}</p>}
      {error && <p>error = {JSON.stringify(error)}</p>}
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          await fetch(new URL(props.method.path, baseUrl), {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(request),
          }).then((res) => {
            if (res.ok) {
              return res.json().then(setResponse);
            }
            return res.json().then(setError);
          });
        }}
      >
        {props.method.fields.map((field) => (
          <label key={field}>
            {field}
            <input
              value={request[field]}
              onChange={(e) => {
                setRequest({ ...request, [field]: e.target.value });
              }}
            />
          </label>
        ))}
        <button type="submit">Send</button>
      </form>
    </div>
  );
}

function MethodList(props: { service: Service }) {
  return (
    <div>
      <h2>{props.service.name}</h2>
      <div>
        {props.service.methods.map((method) => (
          <MethodForm
            method={method}
            key={props.service.name + "/" + method.name}
          />
        ))}
      </div>
    </div>
  );
}

function App() {
  return (
    <div>
      {Proto.fromJson(doc).services.map((service) => (
        <MethodList service={service} key={service.name} />
      ))}
    </div>
  );
}

export default App;
