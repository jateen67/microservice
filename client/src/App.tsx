import "./App.css";

function App() {
  const getRoot = async () => {
    const res = await fetch("http://localhost:8080");
    console.log(await res.text());
  };

  const getHello = async () => {
    const res = await fetch("http://localhost:8080/hello");
    console.log(await res.text());
  };

  return (
    <>
      <button onClick={getRoot}>get root</button>
      <button onClick={getHello}>get hello</button>
    </>
  );
}

export default App;
