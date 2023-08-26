import "./App.css";

function App() {
  const getBroker = async () => {
    const body = {
      Content: "Empty request",
    };

    const res = await fetch("http://localhost:8080", {
      method: "POST",
      body: JSON.stringify(body),
    });

    console.log(body);
    console.log(await res.json());
  };

  const getAuthentication = async () => {
    const body = {
      email: "admin@example.com",
      password: "password123",
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const res = await fetch("http://localhost:8080/authentication", {
      method: "POST",
      body: JSON.stringify(body),
      headers: headers,
    });

    console.log(body);
    console.log(await res.json());
  };

  const getLogger = async () => {
    const body = {
      name: "activity",
      data: "some kind of data",
    };

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const res = await fetch("http://localhost:8080/logger", {
      method: "POST",
      body: JSON.stringify(body),
      headers: headers,
    });

    console.log(body);
    console.log(await res.json());
  };

  return (
    <>
      <button onClick={getBroker}>Broker Service</button>
      <button onClick={getAuthentication}>Authentication Service</button>
      <button onClick={getLogger}>Logger Service</button>
    </>
  );
}

export default App;
