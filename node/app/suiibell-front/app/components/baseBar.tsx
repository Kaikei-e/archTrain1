import { ReactNode } from 'react';
import {
  Box,
  CloseButton,
  Flex,
  Icon,
  useColorModeValue,
  Link,
  Drawer,
  DrawerContent,
  Text,
  useDisclosure,
  BoxProps,
  FlexProps,
  Button,
  DrawerBody,
  DrawerCloseButton,
  DrawerOverlay,
  Spacer,
} from '@chakra-ui/react';
import {
  FiHome,
  FiTrendingUp,
  FiSettings,
  FiEdit,
  FiMenu,
  FiLogOut
} from 'react-icons/fi';
import { IconType } from 'react-icons';
import { ReactText } from 'react';
import { ActionFunction, LoaderFunction, redirect } from '@remix-run/node';
import authenticator from '~/services/auth.server';
import { Form } from '@remix-run/react';

interface LinkItemProps {
  name: string;
  icon: IconType;
  path: string;
}
const LinkItems: Array<LinkItemProps> = [
  { name: 'Home', icon: FiHome, path: "/" },
  { name: 'Trending', icon: FiTrendingUp, path: "trending" },
  { name: 'Logging', icon: FiEdit, path: "logging" },
  { name: 'Settings', icon: FiSettings, path: "settings" },
];

export let loader: LoaderFunction = async ({ request }) => {
  return await authenticator.isAuthenticated(request, {
    failureRedirect: "/login",
  });
};

export const action: ActionFunction = async ({ request }) => {
  await authenticator.logout(request, { redirectTo: "/login" });
};


export default function BaseBar({ children }: { children: ReactNode }) {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <Box>
      <Box h={"6vh"} minW={"100vw"} bgColor={"blue.200"}>
        <Flex bgColor="blue.200">
          <Button ml={"10px"} mt={"5px"} bgColor="blue.200">
            <FiMenu onClick={onOpen} size={"xs"} />
            <Text m={"10px"} fontFamily={"mono"} fontWeight="light" fontSize="3xl">SuiiBell</Text>
          </Button>
          <Spacer />
          <Form method="post">
            <button>
              <Button mr={"10px"} mt={"10px"} bgColor="blue.200" size={"sm"}>
                <FiLogOut />
              </Button>
            </button>
          </Form>

        </Flex>
        <Drawer
          isOpen={isOpen}
          placement="left"
          onClose={onClose}
        >
          <DrawerOverlay />
          <DrawerContent>
            <DrawerCloseButton />
            <DrawerBody bgColor={"blue.200"}>
              <SidebarContent onClose={onClose} />
            </DrawerBody>

          </DrawerContent>
        </Drawer>

      </Box>
    </Box>

  );
}

interface SidebarProps extends BoxProps {
  onClose: () => void;
}

const SidebarContent = ({ onClose, ...rest }: SidebarProps) => {
  return (
    <Box
      bg={useColorModeValue("blue.200", 'gray.900')}
      borderRight="1px"
      borderRightColor={useColorModeValue('blue.200', 'gray.700')}
      w={{ base: 'full', md: 60 }}
      pos="fixed"
      h="full"
      {...rest}>
      <Flex h="20" alignItems="center" mx="8" justifyContent="space-between">
        <Text fontSize="2xl" fontFamily="mono" fontWeight="bold">
          SuiiBell
        </Text>
        <CloseButton display={{ base: 'flex', md: 'none' }} onClick={onClose} />
      </Flex>
      {LinkItems.map((link) => (
        <NavItem key={link.name} path={link.path} icon={link.icon}>
          {link.name}
        </NavItem>
      ))}
    </Box>
  );
};

interface NavItemProps extends FlexProps {
  icon: IconType;
  children: ReactText;
  path: string;
}
const NavItem = ({ icon, path, children, ...rest }: NavItemProps) => {
  return (
    <Link href={path} style={{ textDecoration: 'none' }} _focus={{ boxShadow: 'none' }}>
      <Flex
        align="center"
        p="4"
        mx="4"
        borderRadius="lg"
        role="group"
        cursor="pointer"
        _hover={{
          bg: 'gray.100',
          color: 'black',
        }}
        {...rest}>
        {icon && (
          <Icon
            mr="4"
            fontSize="16"
            _groupHover={{
              color: 'black',
            }}
            as={icon}
          />
        )}
        {children}
      </Flex>
    </Link>
  );
};
