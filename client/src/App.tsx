import "./App.css";

function App() {
  const getRoot = async () => {
    const body = {
      Content: "Empty request",
    };

    const res = await fetch("http://localhost:80", {
      method: "POST",
      body: JSON.stringify(body),
    });

    console.log(body);
    console.log(await res.json());
  };

  return (
    <>
      <button onClick={getRoot}>get root</button>
    </>
  );
}

export default App;
