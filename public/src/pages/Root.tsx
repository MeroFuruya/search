import Search from "../components/Search"

function App() {
  return (
    <>
    <Search onChange={(v) => {console.log(v)}} />
    </>
  )
}

export default App
