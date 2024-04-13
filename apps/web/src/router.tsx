import { loadQuery } from "react-relay"
import { createBrowserRouter } from "react-router-dom"
import rootPageQuery from "./pages/__generated__/rootPageQuery.graphql"
import { RelayEnvironment } from "./relay-env"

const router = createBrowserRouter([
  {
    path: "/",
    lazy: () => import("./pages/root"),
    loader: () => {
      return loadQuery(RelayEnvironment, rootPageQuery, {})
    },
  },
])

export default router
