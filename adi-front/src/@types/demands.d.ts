import { Dayjs } from 'dayjs';

export interface IssuesDetails {
    total: number;
    values: number[];
    jql?: string;
}

export interface DemandsAnalytics {
    overallProgress: number;
    progressPerPeriod: number[];
    createdTotal: number;
    resolvedTotal: number;
    pendingTotal: number;
}

export interface Demands {
    periods: string[];
    created: IssuesDetails;
    resolved: IssuesDetails;
    pending: IssuesDetails;
    analytics: DemandsAnalytics;
}

export interface APIGetCreatedVersusResolvedProps {
    projects: string[];
    period: {
        range: {
            from: Dayjs;
            until: Dayjs;
        };
    };
    orderBy?: {
        field: string;
        direction: 'ASC' | 'DESC';
    }[];
}
