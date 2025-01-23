import supabase from "@/lib/supabase";

export async function POST(request) {
  try {
    const data = await request.json();

    // Get IP from headers (works with Vercel/Cloudflare)
    const rawIP = request.headers.get("x-forwarded-for") || "127.0.0.1";
    const clientIP = rawIP.split(",")[0].trim(); // Handle multiple IPs
    const ipHash = await hashString(clientIP);

    const { error } = await supabase.from("events").insert({
      ...data,
      ip_hash: ipHash,
    });

    console.log("Insert result:", { data, error });

    return new Response(null, {
      status: error ? 500 : 200,
      headers: { "Access-Control-Allow-Origin": "*" },
    });
  } catch (error) {
    console.error("API Error:", error);
    return new Response(null, { status: 500 });
  }
}

async function hashString(str) {
  const encoder = new TextEncoder();
  const data = encoder.encode(str);
  const hashBuffer = await crypto.subtle.digest("SHA-256", data);
  return Array.from(new Uint8Array(hashBuffer))
    .map((b) => b.toString(16).padStart(2, "0"))
    .join("");
}
