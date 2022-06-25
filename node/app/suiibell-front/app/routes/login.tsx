import { Form, useLoaderData } from "@remix-run/react";
import { Box, Button, Flex, Input, Spacer, Text } from "@chakra-ui/react";
import { json, LoaderFunction, ActionFunction } from "@remix-run/node";
import authenticator from "~/services/auth.server";
import { sessionStorage, createUserSession } from "~/services/session.server";

export const loader: LoaderFunction = async ({ request }) => {

  await authenticator.isAuthenticated(request, {
    successRedirect: "/"
  });

  const session = await sessionStorage.getSession(
    request.headers.get("Cookie")
  );

  const error = session.get("sessionErrorKey");
  return json<any>({ error });
};

export const action: ActionFunction = async ({ request, context }) => {
  const resp = await authenticator.authenticate("form", request, {
    successRedirect: "/",
    failureRedirect: "/login",
    throwOnError: true,
    context,
  });

  console.log(resp);
  return resp;
};

export default function LoginPage() {
  const loaderData = useLoaderData();
  console.log(loaderData);
  return (
    <Box
      w={"100vw"}
      h={"100vh"}
      backgroundImage={"url(/img/image1.jpeg)"}
      backgroundSize={"cover"}
      overflow="auto"
      flexDirection="column"
    >
      <Box m={"20px"} flexDirection="column"
        w={"100vw"}
        h={"40vh"}>

        <Flex w={"20%"} style={{ fontFamily: "system-ui, mono", lineHeight: "1.4" }} flexDirection="column">
          <Flex m="5px">
            <Form method="post">
              <Input type="email" name="email" placeholder="user id" required mb={"5px"} color="white"
                _placeholder={{ opacity: 1, color: 'gray.100' }}

              />
              <Input
                type="password"
                name="password"
                placeholder="password"
                autoComplete="current-password"
                mb={"5px"}
                color="white"
                _placeholder={{ opacity: 1, color: 'gray.100' }}
              />
              <button><Button variant="outline" color="white">Sign In</Button></button>
            </Form>
          </Flex>
          <Text bgColor={"white"}>
            {loaderData?.error ? <p>ERROR: {loaderData?.error?.message}</p> : null}
          </Text>
        </Flex>
      </Box>
      <Spacer />
      <Flex align={"end"} justify="end" w={"100vw"} h={"30vh"} flexDirection="column" >
        <Text mr={"20px"} fontSize={"3xl"} fontWeight="semibold" color={"white"}>SuiiBell</Text>
        <Text mr={"20px"} fontSize={"xl"} fontWeight="semibold" color={"white"}>An easy way to manage your balance.</Text>

      </Flex>

    </Box>

  );
}