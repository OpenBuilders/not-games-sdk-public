## not-games-sdk-public@3.0

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install not-games-sdk-public@3.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AchievementsApi* | [**achievementsGet**](docs/AchievementsApi.md#achievementsget) | **GET** /achievements | List achievements
*AchievementsApi* | [**achievementsIdGet**](docs/AchievementsApi.md#achievementsidget) | **GET** /achievements/{id} | Get achievement
*AchievementsApi* | [**achievementsPost**](docs/AchievementsApi.md#achievementspost) | **POST** /achievements | Create achievement
*AuthApi* | [**authGeneratePost**](docs/AuthApi.md#authgeneratepost) | **POST** /auth/generate | Generate new auth
*DropsApi* | [**dropsIdGet**](docs/DropsApi.md#dropsidget) | **GET** /drops/{id} | Check user eligibility for a drop
*DropsApi* | [**dropsIdWhitelistPost**](docs/DropsApi.md#dropsidwhitelistpost) | **POST** /drops/{id}/whitelist | Whitelist users for a drop
*DropsApi* | [**dropsPost**](docs/DropsApi.md#dropspost) | **POST** /drops | Create a new drop
*InventoryApi* | [**inventoryItemsGeneratePost**](docs/InventoryApi.md#inventoryitemsgeneratepost) | **POST** /inventory/items/generate | Generate item
*InventoryApi* | [**inventoryItemsGet**](docs/InventoryApi.md#inventoryitemsget) | **GET** /inventory/items | List items
*InventoryApi* | [**inventoryItemsItemIdConsumePost**](docs/InventoryApi.md#inventoryitemsitemidconsumepost) | **POST** /inventory/items/{item_id}/consume | Consume item
*InventoryApi* | [**inventoryItemsItemIdGet**](docs/InventoryApi.md#inventoryitemsitemidget) | **GET** /inventory/items/{item_id} | Get item
*InventoryApi* | [**inventoryItemsItemIdPut**](docs/InventoryApi.md#inventoryitemsitemidput) | **PUT** /inventory/items/{item_id} | Update item
*InventoryApi* | [**inventoryItemsItemIdTransferPost**](docs/InventoryApi.md#inventoryitemsitemidtransferpost) | **POST** /inventory/items/{item_id}/transfer | Transfer item to another account
*InventoryApi* | [**inventoryItemsItemIdTransferPreflightPost**](docs/InventoryApi.md#inventoryitemsitemidtransferpreflightpost) | **POST** /inventory/items/{item_id}/transfer-preflight | Check transfer item possibility to another account
*InventoryApi* | [**inventoryItemsUnpackPost**](docs/InventoryApi.md#inventoryitemsunpackpost) | **POST** /inventory/items/unpack | Unpack item
*ProfileApi* | [**profileAccountIdGet**](docs/ProfileApi.md#profileaccountidget) | **GET** /profile/{account_id} | Get account
*RegistryApi* | [**registryAchievementsAchievementDefIdDelete**](docs/RegistryApi.md#registryachievementsachievementdefiddelete) | **DELETE** /registry/achievements/{achievement_def_id} | Delete achievement-def
*RegistryApi* | [**registryAchievementsAchievementDefIdGet**](docs/RegistryApi.md#registryachievementsachievementdefidget) | **GET** /registry/achievements/{achievement_def_id} | Get achievement-def
*RegistryApi* | [**registryAchievementsAchievementDefIdPut**](docs/RegistryApi.md#registryachievementsachievementdefidput) | **PUT** /registry/achievements/{achievement_def_id} | Update achievement-def
*RegistryApi* | [**registryAchievementsGet**](docs/RegistryApi.md#registryachievementsget) | **GET** /registry/achievements | List achievement-def
*RegistryApi* | [**registryAchievementsPost**](docs/RegistryApi.md#registryachievementspost) | **POST** /registry/achievements | Create achievement-def
*RegistryApi* | [**registryAppsGet**](docs/RegistryApi.md#registryappsget) | **GET** /registry/apps | List apps
*RegistryApi* | [**registryItemsGet**](docs/RegistryApi.md#registryitemsget) | **GET** /registry/items | List item-def
*RegistryApi* | [**registryItemsItemDefIdDelete**](docs/RegistryApi.md#registryitemsitemdefiddelete) | **DELETE** /registry/items/{item_def_id} | Delete item-def
*RegistryApi* | [**registryItemsItemDefIdGet**](docs/RegistryApi.md#registryitemsitemdefidget) | **GET** /registry/items/{item_def_id} | Get item-def
*RegistryApi* | [**registryItemsItemDefIdPut**](docs/RegistryApi.md#registryitemsitemdefidput) | **PUT** /registry/items/{item_def_id} | Update item-def
*RegistryApi* | [**registryItemsPost**](docs/RegistryApi.md#registryitemspost) | **POST** /registry/items | Create item-def
*SquadsApi* | [**squadsGet**](docs/SquadsApi.md#squadsget) | **GET** /squads | List squads
*SquadsApi* | [**squadsIdGet**](docs/SquadsApi.md#squadsidget) | **GET** /squads/{id} | Get squad by id


### Documentation For Models

 - [AchievementsGet200Response](docs/AchievementsGet200Response.md)
 - [AchievementsIdGet200Response](docs/AchievementsIdGet200Response.md)
 - [AchievementsPost200Response](docs/AchievementsPost200Response.md)
 - [AuthGeneratePost200Response](docs/AuthGeneratePost200Response.md)
 - [DropsIdGet200Response](docs/DropsIdGet200Response.md)
 - [DropsPost200Response](docs/DropsPost200Response.md)
 - [GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest](docs/GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateRequest.md)
 - [GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateResponse](docs/GithubComNotPlatformInternalModuleAuthPresentationAuthorizationHttpModelAuthGenerateResponse.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelConsumeItemRequest.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelGenerateItemRequest.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelInventoryItemResponse](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelInventoryItemResponse.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelListItemsResponse](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelListItemsResponse.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemPreflightResponse](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemPreflightResponse.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelTransferItemRequest.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUnpackItemRequest.md)
 - [GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest](docs/GithubComNotPlatformInternalModuleInventoryPresentationInventoryItemHttpModelUpdateItemRequest.md)
 - [GithubComNotPlatformInternalModuleProfileDomainEntityAccount](docs/GithubComNotPlatformInternalModuleProfileDomainEntityAccount.md)
 - [GithubComNotPlatformInternalModuleProfilePresentationHttpModelAccountBackgroundResponse](docs/GithubComNotPlatformInternalModuleProfilePresentationHttpModelAccountBackgroundResponse.md)
 - [GithubComNotPlatformInternalModuleProfilePresentationHttpModelAccountResponse](docs/GithubComNotPlatformInternalModuleProfilePresentationHttpModelAccountResponse.md)
 - [GithubComNotPlatformInternalModuleRegistryPresentationItemHttpModelItemDefResponse](docs/GithubComNotPlatformInternalModuleRegistryPresentationItemHttpModelItemDefResponse.md)
 - [GithubComNotPlatformInternalModuleRegistryPresentationItemHttpModelListItemDefResponse](docs/GithubComNotPlatformInternalModuleRegistryPresentationItemHttpModelListItemDefResponse.md)
 - [GithubComNotPlatformInternalModuleSharedValueobjectPaginationResponse](docs/GithubComNotPlatformInternalModuleSharedValueobjectPaginationResponse.md)
 - [GithubComNotPlatformInternalServerTemplatesResponseTemplate](docs/GithubComNotPlatformInternalServerTemplatesResponseTemplate.md)
 - [HttpAppOneResponse](docs/HttpAppOneResponse.md)
 - [HttpAppSocial](docs/HttpAppSocial.md)
 - [HttpListAppResponse](docs/HttpListAppResponse.md)
 - [InternalModuleAchievementPresentationHttpCreateAchievementOneResponse](docs/InternalModuleAchievementPresentationHttpCreateAchievementOneResponse.md)
 - [InternalModuleAchievementPresentationHttpCreateAchievementsRequest](docs/InternalModuleAchievementPresentationHttpCreateAchievementsRequest.md)
 - [InternalModuleAchievementPresentationHttpCreateAchievementsResponse](docs/InternalModuleAchievementPresentationHttpCreateAchievementsResponse.md)
 - [InternalModuleAchievementPresentationHttpGetAchievementOneResponse](docs/InternalModuleAchievementPresentationHttpGetAchievementOneResponse.md)
 - [InternalModuleAchievementPresentationHttpListAchievementsResponse](docs/InternalModuleAchievementPresentationHttpListAchievementsResponse.md)
 - [InternalModuleDropsPresentationDropsHttpCreateDropRequest](docs/InternalModuleDropsPresentationDropsHttpCreateDropRequest.md)
 - [InternalModuleDropsPresentationDropsHttpCreateDropResponse](docs/InternalModuleDropsPresentationDropsHttpCreateDropResponse.md)
 - [InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest](docs/InternalModuleDropsPresentationDropsHttpWhitelistDropUserRequest.md)
 - [InternalModuleRegistryPresentationAchievementHttpAchievementDefOneResponse](docs/InternalModuleRegistryPresentationAchievementHttpAchievementDefOneResponse.md)
 - [InternalModuleRegistryPresentationAchievementHttpCreateRequest](docs/InternalModuleRegistryPresentationAchievementHttpCreateRequest.md)
 - [InternalModuleRegistryPresentationAchievementHttpListAchievementDefResponse](docs/InternalModuleRegistryPresentationAchievementHttpListAchievementDefResponse.md)
 - [InternalModuleRegistryPresentationAchievementHttpUpdateRequest](docs/InternalModuleRegistryPresentationAchievementHttpUpdateRequest.md)
 - [InternalModuleRegistryPresentationItemHttpCreateRequest](docs/InternalModuleRegistryPresentationItemHttpCreateRequest.md)
 - [InternalModuleRegistryPresentationItemHttpUpdateRequest](docs/InternalModuleRegistryPresentationItemHttpUpdateRequest.md)
 - [InternalModuleSquadsPresentationSquadsHttpGetSquadOneResponse](docs/InternalModuleSquadsPresentationSquadsHttpGetSquadOneResponse.md)
 - [InternalModuleSquadsPresentationSquadsHttpListSquadsResponse](docs/InternalModuleSquadsPresentationSquadsHttpListSquadsResponse.md)
 - [InventoryItemsGet200Response](docs/InventoryItemsGet200Response.md)
 - [InventoryItemsItemIdGet200Response](docs/InventoryItemsItemIdGet200Response.md)
 - [InventoryItemsItemIdTransferPreflightPost200Response](docs/InventoryItemsItemIdTransferPreflightPost200Response.md)
 - [InventoryItemsUnpackPost200Response](docs/InventoryItemsUnpackPost200Response.md)
 - [ProfileAccountIdGet200Response](docs/ProfileAccountIdGet200Response.md)
 - [RegistryAchievementsAchievementDefIdGet200Response](docs/RegistryAchievementsAchievementDefIdGet200Response.md)
 - [RegistryAchievementsGet200Response](docs/RegistryAchievementsGet200Response.md)
 - [RegistryAppsGet200Response](docs/RegistryAppsGet200Response.md)
 - [RegistryItemsGet200Response](docs/RegistryItemsGet200Response.md)
 - [RegistryItemsItemDefIdGet200Response](docs/RegistryItemsItemDefIdGet200Response.md)
 - [SquadsGet200Response](docs/SquadsGet200Response.md)
 - [SquadsIdGet200Response](docs/SquadsIdGet200Response.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization


Authentication schemes defined for the API:
<a id="JWT"></a>
### JWT

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

