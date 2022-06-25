import { createCookieSessionStorage, redirect } from "@remix-run/node";

export let sessionStorage = createCookieSessionStorage({
  cookie: {
    name: 'suiibell_session',
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
