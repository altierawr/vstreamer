import { PreloadedQuery, graphql, useMutation, usePreloadedQuery } from "react-relay"
import { useLoaderData, useNavigate } from "react-router-dom"
import { rootPageQuery } from "./__generated__/rootPageQuery.graphql"
import { rootPage_createPlaySessionMutation } from "./__generated__/rootPage_createPlaySessionMutation.graphql"

const CreatePlaySessionMutation = graphql`
  mutation rootPage_createPlaySessionMutation($videoID: ID!) {
    createPlaySession(input: { videoID: $videoID }) {
      id
    }
  }
`

const Component = () => {
  const queryRef = useLoaderData() as PreloadedQuery<rootPageQuery>
  const [commitMutation] = useMutation<rootPage_createPlaySessionMutation>(CreatePlaySessionMutation)
  const navigate = useNavigate();

  const data = usePreloadedQuery<rootPageQuery>(
    graphql`
      query rootPageQuery {
        videos {
          id
          path
        }
      }
    `,
    queryRef
  )

  const handleVideoClick = (videoId: string) => {
    commitMutation({
      variables: {
        videoID: videoId
      },
      onCompleted: (resp, errors) => {
        if (errors) {
          console.error("error with video mutation", errors)
          return;
        }

        navigate(`/play/${resp.createPlaySession.id}`)
      }
    })
  }

  return (
    <div className="flex gap-10 p-10">
      {data.videos.map((video) => (
        <div key={video.id} className="">
          <div className="aspect-video bg-gray-300 mb-5 rounded-md cursor-pointer" onClick={() => handleVideoClick(video.id)} />
          <p>{video.path}</p>
        </div>
      ))}
    </div>
  )
}

const ErrorBoundary = () => {
  return <></>
}

export { Component, ErrorBoundary }
