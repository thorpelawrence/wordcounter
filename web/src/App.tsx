import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import WordCounter from "./components/WordCounter";
import "./App.css";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <WordCounter />
      <ReactQueryDevtools initialIsOpen position="bottom-right" />
    </QueryClientProvider>
  );
}

export default App;
