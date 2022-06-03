import { createCookieSessionStorage, redirect } from "@remix-run/node";

export let sessionStorage = createCookieSessionStorage({
  cookie: {
    name: '_session',
    sameSite: 'lax',
    path: '/',
    httpOnly: true,
    secrets: [process.env.SECRET ? process.env.SECRET : "TheDummySecretHogeHoge"],
    secure: process.env.NODE_ENV === 'production',
  },
});

export let { getSession, commitSession, destroySession } = sessionStorage;

export type User = {
  email: string;
  token: string;
};

export async function createUserSession(
  email: string,
  redirectTo: string
) {
  const session = await sessionStorage.getSession();
  session.set("email", email);
  return redirect(redirectTo, {
    headers: {
      "Set-Cookie": await sessionStorage.commitSession(session),
    },
  });
}

export async function getEmail(request: Request) {
  const session = await getUserSession(request);
  const email = session.get("email");
  if (!email || typeof email !== "string") return null;
  return email;
}

export async function getUser(request: Request) {
  const email = await getEmail(request);
  if (typeof email !== "string") {
    return null;
  }

  logout(request);

}

export async function logout(request: Request) {
  const session = await getUserSession(request);
  return redirect("/login", {
    headers: {
      "Set-Cookie": await sessionStorage.destroySession(session),
    },
  });
}


function getUserSession(request: Request) {
  return sessionStorage.getSession(request.headers.get("Cookie"));
}