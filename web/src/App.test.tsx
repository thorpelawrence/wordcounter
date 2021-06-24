import { render, screen } from "@testing-library/react";
import ResultsViewer from "./components/ResultsViewer";

test("renders correct data", async () => {
  render(
    <ResultsViewer
      data={{
        hello: 1,
        world: 2,
        test: 3,
      }}
    />
  );
  expect(screen.getByText(/hello/)).toBeInTheDocument();
  expect(screen.getByText(/test/)).toBeInTheDocument();
  expect(screen.getByRole("table")).toMatchSnapshot();
});
