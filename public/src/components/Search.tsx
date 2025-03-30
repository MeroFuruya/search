import { Component, createSignal } from "solid-js"



const Search: Component<{
  onChange: (v: string) => void
}> = (props) => {
  const [value, setValue] = createSignal("")

  return (
      <div class="flex h-10">
        <input
          class="h-full w-full p-3 border rounded-md"
          type="text"
          placeholder="Search"
          value={value()}
          onInput={(e) => {
            const v = e.currentTarget.value
            setValue(v)
            props.onChange(v)
          }}
        />
      </div>
  )
}
export default Search