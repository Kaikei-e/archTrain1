import { Box, Flex, Text } from "@chakra-ui/react";
import BaseBar from "~/components/baseBar";
import { useLoaderData } from "@remix-run/react";

export default function Index() {
  const data = useLoaderData();
  return (
    <Box minH="100vh" bgColor="gray.100">
      <BaseBar children={undefined} />
      <Flex>
        <Text>Hi, {data?.email}</Text>
      </Flex>
    </Box>
  );
}

