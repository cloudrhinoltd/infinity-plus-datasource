/**
 * This file is based on the original work from Grafana Labs © 2023.
 * Modifications were made by Syncfish Pty Ltd © 2024.
 * The Syncfish Logo and Name are registered Trademarks of Syncfish Pty Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Modifications:
 * - Integrated changes to include `refId` and `url_options` to match the expected type `InfinityQueryWithURLSource<InfinityQueryType>`.
 * - Adjusted `body_content_type` to use a valid value (`'application/json'`) as required by `QueryBodyContentType`.
 * - Added `refId` properties to mock queries to satisfy `DataQuery` type requirements.
 * - Updated the test to ensure compatibility with the modified `InfinityQueryWithURLSource<InfinityQueryType>` type.
 */

import React from 'react';
import { URL } from './query.url';
import { screen, render } from '@testing-library/react';
import { InfinityQueryWithURLSource, InfinityQueryType, InfinityURLOptions, QueryBodyContentType } from 'types';

describe('URL', () => {
  it('should show changed URL', () => {
    // Define a mock InfinityURLOptions for the test
    const mockURLOptions: InfinityURLOptions = {
      method: 'GET',
      headers: [],
      params: [],
      body_content_type: 'application/json' as QueryBodyContentType, // Use a valid value for body_content_type
    };

    const mockQuery1: InfinityQueryWithURLSource<InfinityQueryType> = {
      refId: 'A', // Add required refId property
      url: 'https://example1.com',
      type: 'json',
      source: 'url',
      url_options: mockURLOptions, // Include url_options as required by the expected type
    };

    const props = {
      query: mockQuery1,
      onChange: jest.fn(),
      onRunQuery: jest.fn(),
      onShowUrlOptions: jest.fn(),
    };

    const { rerender } = render(<URL {...props} />);
    expect(screen.getByDisplayValue('https://example1.com')).toBeInTheDocument();

    const mockQuery2: InfinityQueryWithURLSource<InfinityQueryType> = {
      refId: 'B', // Add required refId property
      url: 'https://example2.com',
      type: 'json',
      source: 'url',
      url_options: mockURLOptions, // Include url_options here as well
    };

    rerender(<URL {...props} query={mockQuery2} />);
    expect(screen.getByDisplayValue('https://example2.com')).toBeInTheDocument();
  });
});
