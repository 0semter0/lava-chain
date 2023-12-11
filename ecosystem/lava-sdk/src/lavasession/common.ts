import { Endpoint } from "./consumerTypes";
import { Geolocation, calcGeoLatency } from "./geolocation";
export const AVAILABILITY_PERCENTAGE = 0.1;
export const MAX_ALLOWED_BLOCK_LISTED_SESSION_PER_PROVIDER = 3;
export const MAX_SESSIONS_ALLOWED_PER_PROVIDER = 1000;
export const RELAY_NUMBER_INCREMENT = 1;
export const MAXIMUM_NUMBER_OF_FAILURES_ALLOWED_PER_CONSUMER_SESSION = 3;
export const PERCENTILE_TO_CALCULATE_LATENCY = 0.9;
export const MIN_PROVIDERS_FOR_SYNC = 0.6;
export const DEFAULT_DECIMAL_PRECISION = 18; // same default precision as golang cosmos sdk decimal
export const MAX_CONSECUTIVE_CONNECTION_ATTEMPTS = 5;

export function GetAllProviders(
  allAddresses: Set<string>,
  ignoredProviders: Set<string>
): Array<string> {
  const returnedProviders: string[] = [];

  for (const providerAddress of allAddresses) {
    if (ignoredProviders.has(providerAddress)) {
      // ignored provider, skip it
      continue;
    }
    returnedProviders.push(providerAddress);
  }

  return returnedProviders;
}

export function SortByGeolocations(
  pairingEndpoints: Array<Endpoint>,
  currentGeo: Geolocation
) {
  const latencyToGeo = function (a: Geolocation, b: Geolocation): number {
    return calcGeoLatency(a, b);
  };

  // sort the endpoints by geolocation relevance:
  const lessFunc = function (a: Endpoint, b: Endpoint): Geolocation {
    const latencyA = latencyToGeo(a.geolocation, currentGeo);
    const latencyB = latencyToGeo(b.geolocation, currentGeo);
    return latencyA - latencyB;
  };
  pairingEndpoints.sort(lessFunc);
}
