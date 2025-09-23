/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { consortium_api_hello_v1_HelloRes } from '../models/consortium_api_hello_v1_HelloRes';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class HelloService {
    /**
     * You first hello api
     * @returns consortium_api_hello_v1_HelloRes
     * @throws ApiError
     */
    public static getApiPublicV1Hello(): CancelablePromise<consortium_api_hello_v1_HelloRes> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/public/v1/hello',
        });
    }
}
