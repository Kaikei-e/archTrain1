import { Box, Heading } from "@chakra-ui/react"
import BaseBar from "~/components/baseBar"


function Trending(params:any) {
  return (
    <Box minW={"100vw"} minH={"100vh"} >
      <BaseBar children={undefined}/>
      <Box w={"100vw"} h={"92vh"} overflow={"auto"} bgColor="gray.100">
      <Heading fontWeight={"light"}>Trending: Analyze your balance</Heading>

      </Box>

    </Box>
  )
}

export default Trending