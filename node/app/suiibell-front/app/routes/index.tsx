import { Box, Flex, Text } from "@chakra-ui/react";
import BaseBar from "~/components/baseBar";
import { useLoaderData } from "@remix-run/react";
import { ActionFunction, LoaderFunction } from "@remix-run/node"
import authenticator from "~/services/auth.server";
import { createUserSession } from "~/services/session.server";
import login from "./login";

export default function Index() {
  const data = useLoaderData();
  return (


    <Box minH="100vh" bgColor="gray.100">
      <BaseBar children={undefined} />
      <Flex>
        <Text>Hi, {data?.name}</Text>

      </Flex>
    </Box>
  );
}


export let loader: LoaderFunction = async ({ request }) => {
  return await authenticator.isAuthenticated(request, {
    failureRedirect: "/login",
  });
};

export const action: ActionFunction = async ({ request }) => {
  const user = await login({ username, password });

  await authenticator.logout(request, { redirectTo: "/login" });
  return createUserSession(user.id, redirectTo);
};