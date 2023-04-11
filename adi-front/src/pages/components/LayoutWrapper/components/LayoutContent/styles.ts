import { Layout, styled } from '@adi/react-components';

export const Container = styled(Layout.Content)`
    margin: ${({ theme }) => `calc(${theme.ADIHeaderHeight} + 1rem) 1rem 1rem 1rem`};

    overflow-x: hidden;
`;
