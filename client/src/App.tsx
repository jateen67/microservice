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

  return (
    <>
      <button onClick={getBroker}>Broker Service</button>
    </>
  );
}

export default App;
