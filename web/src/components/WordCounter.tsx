import { useMemo, useState } from "react";
import { useQuery } from "react-query";
import CreatableSelect from "react-select/creatable";
import ResultsViewer from "./ResultsViewer";

const urlToOption = (url: string) => ({
  label: url,
  value: url,
});

const WordCounter = () => {
  const [url, setUrl] = useState(urlToOption("https://bbc.co.uk"));
  const [options, setOptions] = useState(
    ["https://bbc.co.uk", "https://norvig.com/big.txt"].map(urlToOption)
  );

  const apiUrl = useMemo(() => `http://localhost:8080/url/${url.value}`, [url]);

  const { data, status, isFetching, error } = useQuery(["urls", apiUrl], () =>
    fetch(apiUrl).then((res) => res.json())
  );

  return (
    <>
      <CreatableSelect
        value={url}
        options={options}
        onChange={(newValue: any) => setUrl(newValue)}
        onCreateOption={(value) => {
          setOptions([...options, urlToOption(value)]);
          setUrl(urlToOption(value));
        }}
        isLoading={status === "loading"}
      />
      {status === "error" && <h2>Error: {JSON.stringify(error)}</h2>}
      {isFetching ? <h2>Updating {url.value}...</h2> : <h2>{url.value}</h2>}
      {status === "success" && <ResultsViewer data={data} />}
    </>
  );
};

export default WordCounter;
