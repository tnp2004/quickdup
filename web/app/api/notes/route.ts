import { NextRequest } from "next/server";

export async function GET(req: NextRequest) {
    const searchParams = req.nextUrl.searchParams
    const code = searchParams.get("code")
    if (!code) {
        return new Response(JSON.stringify({message: "code can't be empty"}))
    }
    const data = await(await fetch(`${process.env.SERVER_URL}/api/v1/notes/${code}`, {
        method: "GET",
    })).json()
    return new Response(JSON.stringify(data), {headers:{"Content-type": "application/json"}})
}