export async function POST(req: Request, res: Response) {
    const reqData = await req.json()
    const data = await(await fetch(`${process.env.SERVER_URL}/api/v1/notes/`, {
        method: "POST",
        body: JSON.stringify(reqData),
        headers: {
            "Content-type": "application/json"
        },
    })).json()

    return new Response(JSON.stringify(data), {headers: {"Content-type": "application/json"}})
}