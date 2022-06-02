import { Box, Heading } from "@chakra-ui/react"
import BaseBar from "~/components/baseBar"


function Trending(params:any) {
  return (
    <Box minW={"100vw"} minH={"100vh"}>
      <BaseBar children={undefined}/>
      <Heading>Trending: Analyze your balance</Heading>
    </Box>
  )
}

export default Trending