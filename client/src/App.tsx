import "./App.css";

function App() {
  const getRoot = async () => {
    const res = await fetch("http://localhost:80", {
      method: "POST",
    });
    console.log(await res.json());
  };

  return (
    <>
      <button onClick={getRoot}>get root</button>
    </>
  );
}

export default App;
