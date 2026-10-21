import { SpaceflightNewsEntityBase } from '../SpaceflightNewsEntityBase';
import type { SpaceflightNewsSDK } from '../SpaceflightNewsSDK';
import type { Control } from '../types';
import type { Info, InfoLoadMatch, InfoListMatch } from '../SpaceflightNewsTypes';
declare class InfoEntity extends SpaceflightNewsEntityBase<Info> {
    constructor(client: SpaceflightNewsSDK, entopts: any);
    make(this: InfoEntity): InfoEntity;
    load(this: any, reqmatch?: InfoLoadMatch, ctrl?: Control): Promise<InfoEntity>;
    list(this: any, reqmatch?: InfoListMatch, ctrl?: Control): Promise<InfoEntity[]>;
}
export { InfoEntity };
