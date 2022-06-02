import { Box, Flex, Heading } from "@chakra-ui/react"
import BaseBar from "~/components/baseBar"


function Settings(params: any) {
  return (
    <Box minW={"100vw"} minH={"100vh"} >
      <BaseBar children={undefined} />
      <Box w={"100vw"} h={"92vh"} overflow={"auto"} bgColor="gray.100">
        <Flex m={10}>

          <Heading fontWeight={"light"}>Setting: Manage your account.</Heading>
        </Flex>
      </Box>

    </Box>
  )
}

export default Settings