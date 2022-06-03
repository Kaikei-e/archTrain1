import { Box, Flex, Text } from "@chakra-ui/react";
import BaseBar from "~/components/baseBar";
import { useLoaderData } from "@remix-run/react";
import { ActionFunction, LoaderFunction } from "@remix-run/node";
import authenticator from "~/services/auth.server";


export let loader: LoaderFunction = async ({ request }) => {
  return await authenticator.isAuthenticated(request, {
    failureRedirect: "/login",
  });
};

export const action: ActionFunction = async ({ request }) => {
  await authenticator.logout(request, { redirectTo: "/login" });
};

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

