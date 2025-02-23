// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"

	"github.com/okteto/okteto/pkg/types"
)

// FakePreviewsClient mocks the previews interface
type FakePreviewsClient struct {
	response *FakePreviewResponse
}

type FakePreviewResponse struct {
	Preview        *types.PreviewResponse
	PreviewList    []types.Preview
	ResourceStatus map[string]string

	ErrList          error
	ErrDeployPreview error

	ErrResources error
}

// NewFakePreviewClient returns a new fake preview client
func NewFakePreviewClient(response *FakePreviewResponse) *FakePreviewsClient {
	return &FakePreviewsClient{
		response: response,
	}
}

// List list namespaces
func (c *FakePreviewsClient) List(_ context.Context) ([]types.Preview, error) {
	return c.response.PreviewList, c.response.ErrList
}

// DeployPreview deploys a preview
func (c *FakePreviewsClient) DeployPreview(_ context.Context, _, _, _, _, _, _ string, _ []types.Variable) (*types.PreviewResponse, error) {
	return c.response.Preview, c.response.ErrDeployPreview
}

// GetResourcesStatusFromPreview gets resources from a fake preview
func (c *FakePreviewsClient) GetResourcesStatusFromPreview(_ context.Context, _, _ string) (map[string]string, error) {
	return c.response.ResourceStatus, c.response.ErrResources
}
