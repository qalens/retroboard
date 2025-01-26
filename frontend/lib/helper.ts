export function getBaseURL() {
    const isServer = typeof window === 'undefined';
    if (isServer) {
      return process.env.BASE_URL+"/api"
    } else {
      // Client-side (Browser)
      return process.env.NEXT_PUBLIC_BASE_URL+"/api"; // Use relative URLs for client-side
    }
  }