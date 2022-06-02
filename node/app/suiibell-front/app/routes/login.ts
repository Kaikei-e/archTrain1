import { json, LoaderFunction, ActionFunction } from "@remix-run/node";
import authenticator from "~/services/auth.server";

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
  // call my authenticator
  const resp = await authenticator.authenticate("form", request, {
    successRedirect: "/",
    failureRedirect: "/login",
    throwOnError: true,
    context,
  });
  console.log(resp);
  return resp;
};