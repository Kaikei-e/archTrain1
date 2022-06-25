import { createCookieSessionStorage, redirect } from "@remix-run/node";

export let sessionStorage = createCookieSessionStorage({
  cookie: {
    name: 'suiibell_session',
    sameSite: 'lax',
    path: '/',
    httpOnly: true,
    secrets: [process.env.SECRET ? process.env.SECRET : "TheDummySecretHogeHoge"],
    secure: process.env.NODE_ENV === 'production',
    maxAge: 604_800, // 1 week
    expires: new Date(Date.now() + 1000 * 60 * 60 * 24 * 1), // 1 day

  },
});

export let { getSession, commitSession, destroySession } = sessionStorage;

export type User = {
  email: string;
  jwtoken: string;
};

export async function createUserSession(
  user: User,
  redirectTo: string
) {
  const session = await sessionStorage.getSession();
  session.set("email", user.email);
  session.set("jwtoken", user.jwtoken);
  return redirect(redirectTo, {
    headers: {
      "Set-Cookie": await sessionStorage.commitSession(session),
    },

  });
}
