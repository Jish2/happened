/**
 * Generated by orval v7.2.0 🍺
 * Do not edit manually.
 * My API
 * OpenAPI spec version: 1.0.0
 */
import axios from 'axios'
import type {
  AxiosRequestConfig,
  AxiosResponse
} from 'axios'
export interface GreetingOutputBody {
  /** A URL to the JSON Schema for this object. */
  readonly $schema?: string;
  /** Greeting message */
  message: string;
}

export interface ErrorDetail {
  /** Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id' */
  location?: string;
  /** Error message text */
  message?: string;
  /** The value at the given location */
  value?: unknown;
}

/**
 * Optional list of individual error details
 */
export type ErrorModelErrors = ErrorDetail[] | null;

export interface ErrorModel {
  /** A URL to the JSON Schema for this object. */
  readonly $schema?: string;
  /** A human-readable explanation specific to this occurrence of the problem. */
  detail?: string;
  /** Optional list of individual error details */
  errors?: ErrorModelErrors;
  /** A URI reference that identifies the specific occurrence of the problem. */
  instance?: string;
  /** HTTP status code */
  status?: number;
  /** A short, human-readable summary of the problem type. This value should not change between occurrences of the error. */
  title?: string;
  /** A URI reference to human-readable documentation for the error. */
  type?: string;
}





  /**
 * @summary Get greeting by name
 */
export const getGreetingByName = <TData = AxiosResponse<GreetingOutputBody>>(
    name: string, options?: AxiosRequestConfig
 ): Promise<TData> => {
    return axios.get(
      `/greeting/${name}`,options
    );
  }

export type GetGreetingByNameResult = AxiosResponse<GreetingOutputBody>