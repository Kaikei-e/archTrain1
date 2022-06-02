import { Form, useLoaderData } from "@remix-run/react";
import { Box, Button, Flex, Input, Spacer, Text } from "@chakra-ui/react";

export default function LoginPage() {
  // if i got an error it will come back with the loader data
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
              <Input type="email" name="email" placeholder="email" required mb={"5px"} color="white"
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
              <Button variant="outline" color="white">Sign In</Button>
            </Form>
          </Flex>
          <div>
            {loaderData?.error ? <p>ERROR: {loaderData?.error?.message}</p> : null}
          </div>
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